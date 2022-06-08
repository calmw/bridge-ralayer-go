// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.4;

import "./interface/IBridge.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";

contract Bridge is IBridge, AccessControlEnumerable {
    using ECDSA for bytes32;
    using Address for address;

    // AccessControl, Validator role
    bytes32 public constant VALIDATOR_ROLE = keccak256("VALIDATOR");

    // AccessControl, Manager role
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER");

    // Signature threshold, Minimum number of signatures
    uint32 public signatureThreshold;

    // Current chainID
    uint32 public chainId;

    // Current chain nonce
    uint256 public localNonce;

    // Config resource
    struct ConfigResource {
        uint256 fee;
        mapping(address => bool) blackList; // Address state. Whether an address can cross the bridge
    }

    /*
     * Chain config resource , multiple settings can be configured
     * resourceId => ConfigResource
     */
    mapping(bytes32 => ConfigResource) public _configResource;

    // Record state
    enum RecordState {
        Call,
        Confirmed
    }

    // CallRecord struct
    struct CallRecord {
        bytes32 resourceID;
        address caller;
        uint32 targetChainId;
        address target;
        bytes data;
        RecordState recordState;
    }

    /*
     * Call records
     * messageId => CallRecord
     */
    mapping(bytes32 => CallRecord) public callRecords;

    /*
     * Execution record
     * messageId => bool
     */
    mapping(bytes32 => bool) public executionRecord;

    /*
     * Confirmed record
     * messageId => bool
     */
    mapping(bytes32 => bool) public confirmedRecord;

    // Initializes Bridge
    constructor(uint8 _chainId, uint32 _signatureThreshold) {
        chainId = _chainId;
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        signatureThreshold = _signatureThreshold;
    }

    // Set chainId by admin
    function adminSetChainId(uint32 _chainId)
        public
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        chainId = _chainId;
    }

    // Set signature threshold by admin
    function adminSetSignatureThreshold(uint32 _signatureThreshold)
        public
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        signatureThreshold = _signatureThreshold;
    }

    /*
        @notice Sets a new config resource
        @param resourceID resourceID is used to get configuration and other information
        @param _fee callRemote cost.
     */
    function adminSetConfigResource(
        bytes32 resourceID,
        uint256 _fee,
        address[] memory blackList
    ) external {
        require(
            hasRole(MANAGER_ROLE, msg.sender) ||
                hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            "No access "
        );

        require(_fee >= 0, "_fee must be Greater than or equal to 0 ");
        _configResource[resourceID].fee = _fee;

        for (uint256 i = 0; i < blackList.length; i++) {
            _configResource[resourceID].blackList[blackList[i]] = true;
        }
    }

    // Check whether the address is allowed to callRemote
    function _isAllowedToCall(bytes32 resourceID, address _address)
        private
        view
        returns (bool)
    {
        return !_configResource[resourceID].blackList[_address];
    }

    // Cross chain call
    function callRemote(
        bytes32 resourceID,
        uint32 targetChainId,
        address target,
        bytes calldata data
    ) external payable override returns (bytes32) {
        // Check fee
        require(
            msg.value >= _configResource[resourceID].fee,
            "Incorrect fee supplied"
        );

        // Check whether the msg.sender is allowed to call bridge
        require(_isAllowedToCall(resourceID, msg.sender), "Deny of service");

        bytes32 dataHash = keccak256(data);
        bytes32 messageId = keccak256(
            abi.encode(
                chainId,
                resourceID,
                localNonce,
                targetChainId,
                target,
                dataHash
            )
        );

        CallRecord memory record = CallRecord(
            resourceID,
            msg.sender,
            targetChainId,
            target,
            data,
            RecordState.Call
        );

        callRecords[messageId] = record;

        // Emit callRemote event
        emit CallRequest(
            resourceID,
            msg.sender,
            chainId,
            localNonce,
            messageId,
            targetChainId,
            target,
            data
        );

        localNonce++;

        return messageId;
    }

    // Check signatures and call contract
    function execute(
        bytes32 resourceID,
        uint32 sourceChainId,
        uint256 sourceNonce,
        bytes32 messageId,
        uint32 targetChainId,
        address target,
        bytes calldata data,
        bytes[] calldata signatures
    ) external override {
        // Check signatures num
        require(
            signatures.length >= signatureThreshold,
            "Insufficient signatures"
        );

        // Check chainId
        require(
            sourceChainId == chainId || targetChainId == chainId,
            "sourceChainId or targetChainId error"
        );

        // Check if it has been executed
        require(
            (executionRecord[messageId] == false && targetChainId == chainId) ||
                (confirmedRecord[messageId] == false &&
                    sourceChainId == chainId),
            "Can't execute anymore"
        );

        // Verify messageId
        bytes32 dataHash = keccak256(data);
        require(
            keccak256(
                abi.encode(
                    sourceChainId,
                    resourceID,
                    sourceNonce,
                    targetChainId,
                    target,
                    dataHash
                )
            ) == messageId,
            "Can't verify messageId"
        );

        // Check signatures
        for (uint256 i = 0; i != signatures.length; i++) {
            address caller = messageId.toEthSignedMessageHash().recover(
                signatures[i]
            );
            require(hasRole(VALIDATOR_ROLE, caller), "Invalid signature");
        }

        // Determine whether to execute or confirm
        if (targetChainId == chainId) {
            // Call contract
            target.functionCall(data);

            // Change executionRecord state
            executionRecord[messageId] = true;

            // Emit ConfirmedRequest event
            emit ConfirmedRequest(
                resourceID,
                address(this),
                sourceChainId,
                sourceNonce,
                messageId,
                targetChainId,
                target,
                data
            );
        } else if (sourceChainId == chainId) {
            // Change callRecord state,
            callRecords[messageId].recordState = RecordState.Confirmed;

            // Change confirmedRecord state
            confirmedRecord[messageId] = true;
        }
    }
}

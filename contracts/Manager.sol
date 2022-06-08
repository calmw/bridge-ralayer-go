// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.4;

import "./interface/IManager.sol";
import "@openzeppelin/contracts/utils/math/SafeCast.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";

contract Manager is IManager, AccessControlEnumerable {
    using ECDSA for bytes32;

    // AccessControl, validator role
    bytes32 public constant VALIDATOR_ROLE = keccak256("VALIDATOR");

    // AccessControl, manager role
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER");

    // Vote threshold
    uint32 public voteThreshold;

    // Expired block num threshold
    uint256 public ExpiredBlockNum;

    // Nonce for all chains
    mapping(uint32 => uint256) public nonceMap;

    // ResourceConfig option
    enum RemoteCallType {
        Manual,
        Automatic
    }

    // Resource configuration
    struct ResourceConfig {
        RemoteCallType remoteCallType;
    }

    /*
     * config for resourceId
     * resourceId => ResourceConfig
     */
    mapping(bytes32 => ResourceConfig) public chainConfigResource;

    // Vote struct
    struct VoteRecord {
        bytes32 resourceID;
        VoteStatus voteStatus;
        uint256 startBlock; // Block number of voting initiation
        uint32 sourceChainId;
        uint256 sourceNonce;
        uint32 targetChainId;
        bytes32 dataHash;
        bytes[] signatures;
    }

    // Vote records
    mapping(bytes32 => VoteRecord) public voteRecords;

    // Initializes Manager
    constructor(uint32 _voteThreshold, uint256 _expiredBlockNum) {
        ExpiredBlockNum = _expiredBlockNum;
        voteThreshold = _voteThreshold;
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    // Set vote threshold by admin
    function adminSetVoteThreshold(uint32 _voteThreshold)
        public
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        voteThreshold = _voteThreshold;
    }

    // Sets a new config resource
    function adminSetConfigResource(bytes32 resourceID, bool automaticCall)
        external
    {
        require(
            hasRole(MANAGER_ROLE, msg.sender) ||
                hasRole(DEFAULT_ADMIN_ROLE, msg.sender),
            "No access "
        );

        // AutomaticCall
        if (automaticCall) {
            chainConfigResource[resourceID].remoteCallType = RemoteCallType
                .Automatic;
        } else {
            chainConfigResource[resourceID].remoteCallType = RemoteCallType
                .Manual;
        }
    }

    // Get chain config by resourceID
    function getChainConfig(bytes32 resourceID)
        external
        view
        returns (ResourceConfig memory)
    {
        return chainConfigResource[resourceID];
    }

    // Vote
    function vote(
        bytes32 _resourceID,
        bytes32 messageId,
        uint32 sourceChainId,
        uint256 sourceNonce,
        uint32 targetChainId,
        address target,
        bytes32 dataHash,
        bytes calldata signature
    ) external override onlyRole(VALIDATOR_ROLE) {
        // Verify messageId.
        require(
            keccak256(
                abi.encode(
                    sourceChainId,
                    _resourceID,
                    sourceNonce,
                    targetChainId,
                    target,
                    dataHash
                )
            ) == messageId,
            "Can't verify messageId"
        );

        // Check caller
        address caller = messageId.toEthSignedMessageHash().recover(signature);
        require(caller == msg.sender, "signature must be signed by caller");

        VoteRecord storage voteRecord = voteRecords[messageId];

        // Check voting state
        require(
            voteRecord.voteStatus == VoteStatus.Inactive ||
                voteRecord.voteStatus == VoteStatus.Active,
            "The vote already passed/cancelled"
        );

        // Check whether the validator has voted
        require(!_hasVoted(messageId, caller), "validator already voted");

        // Init vote records
        if (voteRecord.voteStatus == VoteStatus.Inactive) {
            voteRecord.resourceID = _resourceID;
            voteRecord.sourceChainId = sourceChainId;
            voteRecord.targetChainId = targetChainId;
            voteRecord.voteStatus = VoteStatus.Active;
            voteRecord.startBlock = block.number;
        }

        // If the number of blocks that has passed since this vote was submitted exceeds the ExpiredBlockNum threshold set
        require(
            (block.number - voteRecord.startBlock) < ExpiredBlockNum,
            "Voting expires"
        );

        voteRecord.sourceNonce = sourceNonce;
        voteRecord.dataHash = dataHash;
        voteRecord.signatures.push(signature);

        if (voteRecord.signatures.length >= voteThreshold) {
            // Change vote state
            voteRecord.voteStatus = VoteStatus.Passed;

            // Save source chain nonce
            nonceMap[sourceChainId] = sourceNonce;

            // Emit an event
            emit SignatureCollected(
                _resourceID,
                VoteStatus.Passed,
                sourceChainId,
                sourceNonce,
                targetChainId,
                target,
                messageId,
                voteRecord.dataHash,
                voteRecord.signatures
            );
        }
    }

    // Check if validator has voted on message.
    function _hasVoted(bytes32 messageId, address validator)
        private
        view
        returns (bool)
    {
        VoteRecord storage voteRecord = voteRecords[messageId];
        if (uint256(voteRecord.signatures.length) <= 0) {
            return false;
        }

        for (uint256 i = 0; i != voteRecord.signatures.length; i++) {
            address caller = messageId.toEthSignedMessageHash().recover(
                voteRecord.signatures[i]
            );
            if (caller == validator) {
                return true;
            }
        }

        return false;
    }

    // Check if validator has voted on message.
    function hasVotedOnMessage(bytes32 messageId) external view returns (bool) {
        require(hasRole(VALIDATOR_ROLE, msg.sender), "No access");
        return _hasVoted(messageId, msg.sender);
    }
}

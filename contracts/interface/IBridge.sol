// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.4;

/**
    @title Interface for Bridge contract.
 */
interface IBridge {
    function callRemote(
        bytes32 resourceID,
        uint32 targetChainId,
        address target,
        bytes calldata data
    ) external payable returns (bytes32);

    event CallRequest(
        bytes32 resourceID,
        address caller,
        uint32 sourceChainId,
        uint256 sourceNonce,
        bytes32 indexed messageId,
        uint32 indexed targetChainId,
        address indexed target,
        bytes data
    );

    event ConfirmedRequest(
        bytes32 resourceID,
        address caller,
        uint32 sourceChainId,
        uint256 sourceNonce,
        bytes32 indexed messageId,
        uint32 indexed targetChainId,
        address indexed target,
        bytes data
    );

    function execute(
        bytes32 resourceID,
        uint32 sourceChainId,
        uint256 sourceNonce,
        bytes32 messageId,
        uint32 targetChainId,
        address target,
        bytes calldata data,
        bytes[] calldata signatures
    ) external;
}

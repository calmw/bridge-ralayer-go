// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity ^0.8.4;

/**
    @title Interface for Bridge contract.
 */
interface IManager {
    enum VoteStatus {
        Inactive,
        Active,
        Passed,
        Cancelled
    }
    event SignatureCollected(
        bytes32 resourceID,
        VoteStatus voteStatus,
        uint32 sourceChainId,
        uint256 sourceNonce,
        uint32 targetChainId,
        address target,
        bytes32 indexed messageId,
        bytes32 indexed dataHash,
        bytes[] signatures
    );

    function vote(
        bytes32 resourceID,
        bytes32 messageId,
        uint32 sourceChainId,
        uint256 sourceNonce,
        uint32 targetChainId,
        address target,
        bytes32 dataHash,
        bytes calldata signature
    ) external;
}

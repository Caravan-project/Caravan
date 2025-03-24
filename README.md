# Caravan: Incentive-Driven Account Migration via Transaction Aggregation in Sharded Blockchain
## Description
Caravan is an optimized account migration scheme based on Fine-tuned Lock [INFOCOM'24]. In brief: (1) Caravan introduces a transaction aggregation mechanism to efficiently process withdrawal transactions associated with migrating accounts, while ensuring security through a modified multi-level Merkle tree structure; (2) Caravan proposes an incentive-driven priority mechanism for migration transactions. By increasing the revenue generated from these transactions, it incentivizes miners to prioritize them, thereby accelerating the migration process.

Caravan is currently under submission, with its preprint available on ePrint. For a detailed view of Caravan’s design, refer to: xxx.

Note that our experimental prototype of Caravan is built on the open-source blockchain testbed BlockEmulator, specifically the fine-tune-lock branch. 
The HuangLab team at Sun Yat-sen University has developed comprehensive documentation on this open-sourced testbed. 
For further knowledge on sharded blockchains and account migration, refer to: 

**BlockEmulator**. https://github.com/HuangLab-SYSU/block-emulator/tree/main 

**fine-tune-lock branch.** https://github.com/HuangLab-SYSU/block-emulator/tree/Fine-tune-lock.




## Run a node
Running a node in Caravan is the same as Fine-tuned Lock.

Here is an example:
```
go run main.go -S 4 -f 1 -s S1 -n N1 -t 20W.csv
```


## Batch running
We provide an example of a batched startup for a system consisting of four shards, each containing seven nodes, in `batch_run.bat`.
For systems of other sizes, users can write their own batch files for batched startup as needed.

## Raw experimental data

In our experiments, we set the transaction arrival rates to 400, 500, 1000, and 2000 TXs/sec and configure the account migration intervals to 25 and 50 blocks. Each migration is executed by invoking algorithm CLPA [1]. The network size is set to 2 and 4 shards, with each shard containing 7 nodes. A block is set to limit of 2,000 transactions. The capacity of a block is set to 2000 TXs. The raw experimental data of three schemes (LB-Chain [2], Fine-tuned Lock [3] and Caravan) is given in `Raw_experimental_data.7z`.

[1] Li C, Huang H, Zhao Y, et al. Achieving scalability and load balance across blockchain shards for state sharding[C]//2022 41st International Symposium on Reliable Distributed Systems (SRDS). IEEE, 2022: 284-294.

[2] Li M, Wang W, Zhang J. LB-Chain: Load-balanced and low-latency blockchain sharding via account migration[J]. IEEE Transactions on Parallel and Distributed Systems, 2023, 34(10): 2797-2810.

[3] Huang H, Lin Y, Zheng Z. Account Migration across Blockchain Shards using Fine-tuned Lock Mechanism[C]//IEEE INFOCOM 2024-IEEE Conference on Computer Communications. IEEE, 2024: 271-280.

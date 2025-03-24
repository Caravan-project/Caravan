# Description

Caravan is an optimized account migration scheme based on Fine-tuned Lock [INFOCOM'24]. In brief: (1) Caravan introduces a transaction aggregation mechanism to efficiently process withdrawal transactions associated with migrating accounts, while ensuring security through a modified multi-level Merkle tree structure; (2) Caravan proposes an incentive-driven priority mechanism for migration transactions. By increasing the revenue generated from these transactions, it incentivizes miners to prioritize them, thereby accelerating the migration process.

Caravan is currently under submission, with its preprint available on ePrint. For a detailed view of Caravan’s design, refer to: xxx.

Note that our experimental prototype of Caravan is built on the open-source blockchain testbed BlockEmulator, specifically the fine-tune-lock branch. For foundational knowledge on sharded blockchains and account migration, refer to: https://github.com/HuangLab-SYSU/block-emulator/tree/main and https://github.com/HuangLab-SYSU/block-emulator/tree/Fine-tune-lock. The HuangLab team at Sun Yat-sen University has developed comprehensive documentation based on this work.


## Run a node
Running a node in Caravan is the same as Fine-tuned Lock
Here is an example:
```
go run main.go -S 4 -f 1 -s S1 -n N1 -t 20W.csv
```


## Batch running
xxxxxxx


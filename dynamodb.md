# Amazon DynamoDB Cheatsheet

- DynamoDB Core Components:
  - Tables (a collection of items)
  - Items (a collection of attributes)
  - Attributes
  - Primary Keys (to uniquely identify each item in a table)
  - Secondary Indexes (to provide more querying flexibility)
  - DynamoDB Streams (to capture data modification events in tables)

- Attributes:
  - Scalar Attributes
    - e.g. Strings, Numbers, Binaries
  - Nested Attributes
    - Could be nested up to 32 levels deep

- Primary Keys
  - Each item in the table must have the primary key attribute(s)
  - Consist of 1 or 2 attribue(s):
    - Partition Key (1 Attribute)
    - Partition Key and Sort Key (2 Attributes)
  - Each primary key attribute must be either string, number, or binary
  - Other Names:
    - Partition Key = Hash Attribute
    - Sort Key = Range Attribute

- Secondary Indexes
  - You can create one or more secondary indexes on a table.
  - A secondary index lets you query the data in the table using an alternate key, in addition to queries against the primary key.
  - Kinds:
    - Global secondary index
      - An index with a partition key and sort key that can be different from those on the table.
    - Local secondary index
      - An index that has the same partition key as the table, but a different sort key.
  - You can define up to 5 global secondary indexes and 5 local secondary indexes per table.

- DynamoDB Streams
  - DynamoDB Streams is an optional feature that captures data modification events in DynamoDB tables.

- The DynamoDB API
  - Control Plane
    - CreateTable
    - DescribeTable
    - ListTables
    - UpdateTable
    - DeleteTable
  - Data Plane
    - Creating Data
      - PutItem
      - BatchWriteItem (max. 25 items)
    - Reading Data
      - GetItem
      - BatchGetItem (max. 100 items)
      - Query
      - Scan
    - Updating Data
      - UpdateItem
    - Deleting Data
      - DeleteItem
      - BatchWriteItem (max. 25 items)
  - DynamoDB Streams
    - ListStreams
    - DescribeStream
    - GetShardIterator
    - GetRecords

- Data Types
  - Scalar Types
    - String
    - Number
    - Binary
    - Boolean
    - Null
  - Document Types
    - List
    - Map
  - Set Types
    - StringSet
    - NumberSet
    - BinarySet

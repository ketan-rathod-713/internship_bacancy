# Problem Statement

In schema as server_name is not a unique, so we were not able to reference it in forein key, as it requires reference column to be either primary key or foreign key.

Problem solved by making server_name field as unique in gs_info and in server_info tables.

# Setup

- Migrate cli would be good for running this queries Or by simply we can run 3 up files in sequence.

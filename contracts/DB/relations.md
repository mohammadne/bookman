# Book-Store Database Design

## Details

| entities             | records                                                                 |
|----------------------|-------------------------------------------------------------------------|
| members              | id(PK), role_id(FK), contact_informations_id(FK), first_name, last_name |
| roles                | id(PK), name, description                                               |
| contact_informations | id(PK), phone_number, email                                             |
| books                | id(PK), name, publisher_id(FK), language_id(FK)                         |
| publishers           | id(PK), name, location                                                  |
| languages            | id(PK), name                                                            |
| comments             | id(PK), description, books_id(FK), member_id(FK)                        |
| categories           | id(PK), name, description                                               |
| orders               | id(PK), members_id(FK), payments_id(FK), data_time                      |
| payments             | id(PK), amount, status                                                  |
| deliveries           | id(PK), branches_id(FK), orders_id(FK)                                  |
| branches             | id(PK), name                                                            |
| books_orders         | id(PK), books_id(FK), orders_id(FK), count                              |
| books_categories     | id(PK), books_id(FK), categories_id(FK)                                 |
| books_members        | id(PK), books_id(FK), members_id(FK)                                    |

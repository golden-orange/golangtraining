adding database

-setting up postgresql
- connecting postgresql to go
- setting up tables

- using GORM to interact with database


NOTES:
- eventual consistency (sec8.4) - is ok for updates that are not required to be immediate
(eg posts to a blog), unlike updates to a shopping cart or other transactions.

postgresql scales well, up to 1m users easily handled (with proper adjustments)
do not spend time optimizing for problems you don't have
code is meant to be destroyed, it is never written to stand the test of time.
code is always altered and ripped up and rewritten/written over.

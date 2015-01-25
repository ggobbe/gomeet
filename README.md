# GoMeet

We are two expatriates developers in Scotland, therefore we are looking to meet people sharing the same interests as we do but between individuals instead of as a group (as it is possible through the MeetUp platform). Through GoMeet we hope to make that wish possible!
Please note that the current version is a draft/PoC made during and for the GophersGala, we plan to continue working on it to make it something usable for and by all.

Project running on [http://gomeet.cloudapp.net](http://gomeet.cloudapp.net)

## Set up

The only requirement is to have an instance of **MongoDB** running on ***localhost***. There is some Test Data available in the data folder in the file `mgo_users.json` that can be loaded as per the following section (Test Data Import).

## Test Data

### Import
`mongoimport -d gomeet -c users < data\mgo_users.json`

### Export
`mongoexport -d gomeet -c users -o daata\mgo_users.json`

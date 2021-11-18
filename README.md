# Gramaddict Helper
Gramaddict web UI interface for a single device
allow you to see the logs in real time, start and stop the bot all from your browser.

Intended to be used on a single device attached to a controlling machine (e.g. raspberry pi with ADB over USB).

### How to use

Generate binary `make binary`
move it to your `$PATH`

Run it as a service or in a background process (less recommended):
`APP="account_value" GIN_MODE="release" gramaddict_helper`

Cronjob `curl localhost:8080/start` based on your timing needs.

### Environment variables

`ACC` [Mandaotory]    Instagram account name

`GH` [Optional]       Gramaddict installation directory (default to '$HOME/gramaddict')

`FILTER` [Optional]   Used to determine if the process running by filtering the processes (default to 'run.py')

`LOG` [Optional]      Path to gramaddict account log file (default to '$HOME/gramaddict/logs/$INSTAGRAM_ACCOUNT.log

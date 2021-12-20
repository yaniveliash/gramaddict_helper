# Gramaddict Helper
Gramaddict web UI interface for a single device
allow you to see the logs in real time, start and stop the bot all from your browser.

Intended to be used on a single device attached to a controlling machine (e.g. raspberry pi with ADB over USB).

### How to use

Generate binary `make binary`
move it to your `$PATH` or add a new entry in your `$PATH`.

Set it to run on reboot using cronjob (replace env vars):
`@reboot ACC="ACCOUNT_NAME" GH="PATH_TO_GRAMADDICT" gramaddict_helper`

Then, add a cronjob to execute your bot based on your required time (an example for everyday at 07:00 in the morning):
`0 07 * * * curl localhost:8080/start` based on your timing needs.

### Environment variables

`ACC` [Mandaotory]    Instagram account name

`GH` [Optional]       Gramaddict installation directory (default to '$HOME/gramaddict')

`FILTER` [Optional]   Used to determine if the process running by filtering the processes (default to 'run.py')

`LOG` [Optional]      Path to gramaddict account log file (default to '$HOME/gramaddict/logs/$INSTAGRAM_ACCOUNT.log

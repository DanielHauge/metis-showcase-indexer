# Ideas
select github repos:

https://github.com/DanielHauge/goSpace

https://github.com/DanielHauge/plex-folder-soldier

Index-config:{ Url: string, Reindex: every 30 min / on hook updates }

Trigger reindex -> Reindex queue

Reindex queue -> Worker

- Download Files
  - Index code file with (Lines, Language, ProjectId, Filename)
  - Analyse code for smells and add to notification messages.
  - Analyse typical metrics and add to stats
- Accumulate metrics and stats
- Find showcase file and save tabs, stats and metrics to redis


## Issues
- Repo space is not queue, it is stack, it never loops.

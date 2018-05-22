@title[Code Presenting]
# Code
# Presenting

---
@title[Working With Code Blocks]

### Code-Blocks

#### The Basics

![Press Down Key](assets/down-arrow.png)

+++
@title[Sample Block]

```python
from time import localtime

activities = {8: 'Sleeping', 9: 'Commuting', 17: 'Working',
              18: 'Commuting', 20: 'Eating', 22: 'Resting' }

time_now = localtime()
hour = time_now.tm_hour

for activity_time in sorted(activities.keys()):
    if hour < activity_time:
        print activities[activity_time]
        break
else:
    print 'Unknown, AFK or sleeping!'
```
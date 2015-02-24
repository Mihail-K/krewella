krewella
========

IRC notification bot based on HTTP requests

API
---

### `/message/:network/:channel`

#### POST

Requests the post body (as plain text) to be sent to :network's :channel

:channel will automatically have the proper channel name prefix as RFC 1459 
dictates.

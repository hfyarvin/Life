# service application
## application
- /api/applications  --post
- /api/applications  --get
- /api/applications/{id}  --get
- /api/applications/{id}  --put
- /api/applications/{id}  --delete

## application_user
- /api/applications/{id}/users  --post
- /api/applications/{id}/users  --get
- /api/applications/{id}/users/{userID}  --get
- /api/applications/{id}/users/{userID}  --put
- /api/applications/{id}/users/{userID}  --delete

## application-integration
- /api/applications/{id}/integrations/http  --post
- /api/applications/{id}/integrations/http  --get
- /api/applications/{id}/integrations/http  --put
- /api/applications/{id}/integrations/http  --delete
- /api/applications/{id}/integrations  --get  (all configured integrations)

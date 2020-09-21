![plaudern](https://raw.githubusercontent.com/bst27/plaudern/master/website/banner.png)

# This project is work in progress
## TODOS
### MVP
* As an admin I want to approve comments [x]
* As an admin I want to protect my backend with credentials [x]
* As an admin I want my backend to be CSRF secured
* As an admin I want to create a backup / extract my data

### Backlog
* As an admin I want to change comment messages
* As an admin I want to change comment authors
* As an admin I want to change comment thread IDs
* As an admin I want to delete comments
* As an admin I may want to collect the users email address
* As an admin I want to disable/enable comments for selected threads
* As an admin I want to auto-approve new comments via webhook
* As an admin I want to filter/search comments
* As a desktop user I want to have an optimized UI for large screens
* As an admin I want to have a translated webinterface
* As an admin I want to format new comments via webhook

# About
Plaudern is an app to manage comments for static websites.

# Examples
See the [examples](examples) directory for examples.

# Build
To build executables for multiple platforms as well as the Angular based webinterface
you can use the build script at `scripts/build.sh`.

To simplify webinterface development jump into the `web` directory and use
`ng serve --host 0.0.0.0 --disableHostCheck --open` to launch a hot-reloading webserver
which is accessible from other devices, too. If you only want to make it available
locally use `ng serve --open`. The webserver will proxy API requests to `http://localhost:8080`
so make sure you have the backend up and running on this endpoint. You can use
`scripts/develop.sh` which handles this for you.
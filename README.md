# ginmicro
GIN Microservice to read from MQTT broker and present the data on a webpage

## Design of GIN app
Request -> Router Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response

## Router configuration (will be changed after)
- Serve the index page at route / (HTTP GET request), 
- Group user-related routes under the /u route, 
-- Serve the login page at /u/login (HTTP GET request), 
-- Process the login credentials at /u/login (HTTP POST request), 
-- Log out at /u/logout (HTTP GET request), 
-- Serve the registration page at /u/register (HTTP GET request), 
-- Process the registration information at /u/register (HTTP POST request) , 
- Group article related routes under the /article route, 
-- Serve the article creation page at /article/create (HTTP GET request), 
-- Process the submitted article at /article/create (HTTP POST request), and
-- Serve the article page at /article/view/:article_id (HTTP GET request). Take note of the :article_id part in this route. The : at the beginning indicates that this is a dynamic route. This means that :article_id can contain any value and Gin will make this value available in the route handler.

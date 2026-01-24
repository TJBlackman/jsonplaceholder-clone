Read file README.md
Do not proceed until you've read the entire file.

Task:
Add a new route for /api/echo
it returns json { timestamp, message: "ok" }
It takes an optional query parameter of ?message= and if the that's present the response is { timestamp, message: {message} }

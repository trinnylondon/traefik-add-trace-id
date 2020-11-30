# About

This plugin will append a custom header for tracing with a random value if one is not found already in the incoming request.

You can optionally customise this by specifying a custom header name that the plugin will look for in the incoming request (defaults to `X-Trace-Id`) and you can also specify a custom prefix to be added to that header (defaults to `""`).
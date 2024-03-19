# cosmos-experiments
An app-chain that experiments with Cosmos SDK features.

# Ideas and Experiments to Try

## Modules

* An observability module that pushes events, logs, and metrics to something like Datadog


# Developer Notes

## Working in WSL

### Finding the IP address of your WSL container

If you're running `ignite serve` and you're trying to connect to the
faucet on http://0.0.0.0:4500/ then you'll need the IP address of your
WSL container first.

Open a powershell or command prompt session and run `wsl hostname -I`

## Debugging

Install [https://github.com/go-delve/delve/tree/master/Documentation/installation](dlv) and
then use `make debug` to start the server in debug mode.

If you're using VSCode, there's a debug configuration called "Debug Chain" that will attach
to the running process.

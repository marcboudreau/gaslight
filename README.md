# gaslight
Gaslight is a **G**CP **A**utonomous **S**warm setup tool.  This tool is
designed to be executed from the startup script of VM instances managed by
an Instance Group Manager.  The tool uses arguments fed via the startup
script as well as metadata obtained from the metadata endpoint to determine
how the VM instance should join the Docker Swarm, or whether it needs to
initialize a new Docker Swarm.

## How it Works

The **gaslight** tool works for both VM instances intended as manager nodes
and worker nodes.  This flow chart describes the steps and decision points that
the tool follows.

![Flow chart](./flowchart.png)

## Building

## Testing

## Contributing

## Licensing

MIT License.  Refer to the [LICENSE](./LICENSE) file for the specific terms and
conditions.



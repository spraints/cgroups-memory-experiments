# Experiment with memory caps in docker (or cgroups more generally)

Try it out! The child processes will get killed off by the kernel, but the
parent process continues running and gracefully shuts down.

    ./run --count 10 --rate 500m

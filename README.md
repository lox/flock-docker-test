## Hypothesis:

flock() based locking doesn't work between docker containers.

Re: https://github.com/buildkite/agent/pull/1624

## Running

```bash
. bin/activate-hermit
./test.sh
```

If all the containers immediately get locks, then it locks aren't making it out of the containers.

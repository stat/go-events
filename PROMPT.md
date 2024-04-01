# Grid Code Challenge

For this part of the interview, we're asking you to work through a practical software engineering problem.

Our inspiration is https://flightradar24.com. FlighRadar24 displays a live map of airplanes all over the world based on a crowd-sourced dataset.

Our goal now is to build a small feature that would be a part of their overall product implementation.

## Context

The flight tracking backend merges data from two sources primarily:

1. The FAA's official flight database - daily list of flight
2. Our network of crowd-sourced ADS-B receivers - real-time flight locations

Because the network data is crowd-sourced, it receives many duplicate, late or otherwise noisy events.

Let's address this, so that the rest of our system can work with a clean event stream with a **nominal throughput of 1 location per second per active aircraft**.

Any duplicates, fraudulent or late events should be logged for later analysis but otherwise ignored in the live system.

## Scope

Let's build an event cleaning system that consumes a chaotic, noisy event stream of timestamped flight locations and produces a uniform, clean event stream.

The input is random and may include more or less than one event per flight per second. The output should include, on average, one event per second per active flight.

Let's focus this effort on a unit test which generates interesting lists of events and validates the behavior of this event cleaning algorithm. Start with small scale (1-5 hard-coded events), and if you have time, simulate the system at scale (e.g. 1M events over 10k unique flights).

For the sake of this code exercise, let's just focus on application code. You are welcome to use production infrastructure components if that is faster, but it's also ok to substitute databases, etc. with in-memory data-structures.

## Requirements

### Consumer: Raw Events

Saves an event into the system, one at a time.

- Input: tuples of `timestamp, latitude, longitude, station id`
- Output: acknowledgement that event was saved

_Note: This function will only receive events for a single aircraft id, so you do not need to partition the events by aircraft._

Exception Cases:
- TBD

### Producer: Cleaned Events

Returns a reduced list of cleaned, defrauded, de-duped locations since the last call to this API, approx 1 event per second.

- Output: Location Events: `timestamp, latitude, longtidue`

Exception Cases:
- TBD

## Interview Tips

Please avoid spending time on boilerplate. You can feel free to use an in-memory data structure or sqlite instead of a production-grade database or message broker.

You should write this code the way you expect to write production code at Grid. Feel free to use whatever programming environment, database, open source libraries, etc., you'd like.

You should choose the tech stack that you are most comfortable with for production applications today, preferably [Go](https://go.dev). Please do not learn a new or exotic tech stack for this code sample!

Please replace this README with a proper documentation for your solution. You should include information about the production infrastructure setup that is intended. However, you do not need to spend time on the actual deployment.

Please include instructions for other developers to run this project. Ideally, you should include a `Dockerfile`, so that it's easy to run in any environment.

Please submit your solution via pull request on this private repo. To keep this exercise fair for other candidates, we'd prefer you didn't post it open source on github.

We’re not looking for any “right answer” in particular, but rather a practical solution that works overall. Feel free to cut scope for time and focus on an MVP.

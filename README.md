# Dining Philosophers Problem

This Go program demonstrates the classic dining philosophers problem with specific constraints and modifications. The dining philosophers problem is a synchronization and concurrency problem that illustrates challenges in resource sharing among multiple processes.

## Problem Constraints

1. There are 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times.
3. The philosophers pick up the chopsticks in any order.
4. In order to eat, a philosopher must get permission from a host, which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered from 1 through 5.
7. When a philosopher starts eating, it prints "starting to eat \<number\>" on a line by itself.
8. When a philosopher finishes eating, it prints "finishing eating \<number\>" on a line by itself.

## Implementation Details

The program is implemented in Go and consists of the following components:

- `ChopS` struct: Represents a chopstick with its associated mutex.
- `Philo` struct: Represents a philosopher with left and right chopsticks, the number of times eaten, and an ID.
- `hostChanel`: A buffered channel used by the host to limit the number of philosophers eating concurrently.
- `philWaitGroup`: A wait group to ensure all philosophers finish eating before the program exits.

## How to Run

1. Ensure you have Go installed on your system.
2. Save the code to a file, for example, `phil.go`.
3. Open a terminal and navigate to the directory containing the file.
4. Run the program using the command:
    ```bash
    go run phil.go
    ```

## Expected Output

The program output demonstrates the philosophers acquiring and releasing chopsticks, along with the start and finish messages. The philosophers should each eat three times, and the program will terminate when all philosophers have finished eating.


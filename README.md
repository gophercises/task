# Exercise #7: CLI Task Manager

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/task) <!--[![demo](https://img.shields.io/badge/demo-%E2%86%92-yellow.svg?style=for-the-badge)](https://gophercises.com/demos/cyoa/)-->

## Exercise details

In this exercise we are going to be building a CLI tool that can be used to manage your TODOs in the terminal. The basic usage of the tool is going to look roughly like this:

```
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
You have completed the "review talk proposal" task.

$ task list
You have the following tasks:
1. some task description
```

*Note: Lines prefixed with `$` are lines where we type into the terminal, and other lines are output from our program.*

Your final CLI won't need to look exactly like this, but this is what I roughly expect mine to look like. In the bonus section we will also discuss a few extra features we could add, but for now we will stick with the three show above:

- `add` - adds a new task to our list
- `list` - lists all of our incomplete tasks
- `do` - marks a task as complete

In order to build this tool we are going to need to explore a few different topics. Most notably, we will need to:

1. Learn about creating command line interfaces (CLIs)
2. Interact with a database. We will be using BoltDB in this exercise so we can learn about it.
3. Figure out how to store our database file on different operating systems. This will basically boil down to learning about home directories.
4. Exit codes (briefly)
5. And probably more. I'll update this list once the exercise is done.

You are welcome to tackle the problem however you see fit, but below is the order I would recommend to start.

### 1. Build the CLI shell

For building the CLI, I highly recommend using a third party package (library, framework, or whatever you want to call it). You can do this exercise without one, but there are a lot of edge cases you will need to handle on your own and in this case I think it is best to just pick an existing library to use.

There are a lot of CLI libraries, and you can find most of them here: <https://github.com/avelino/awesome-go#command-line>

When I code this exercise I intend to use [spf13/cobra](https://github.com/spf13/cobra). It isn't necessarily better than others out there, but it is one I have used in the past and I know it will serve my needs.

Once you decide on a library, use it to create the original `task` command that displays all your subcommands, and then create stubbed subcommands for each of the actions we discussed above. The actions don't actually have to do anything with a database just yet, but we want to make sure the user typing each individual command will result in a different piece of code running.

For instance, let's say we defined the `task list` command to run the following Go code:

```go
fmt.Println("This is a fake \"list\" command")
```

Then when we used that command with our CLI we should see the following:

```
$ task list
This is a fake "list" command
```

After stubbing out all 3 commands, try to also look at how to parse arguments for the `task do` and `task add` commands.

### 2. Write the BoltDB interactions

After stubbing out your CLI commands, try writing code that will read, add, and delete data in a BoltDB database. You can find more information about using Bolt here: <https://github.com/boltdb/bolt>

*Note: I know many people claim `bolt` is abandoned, but that is inaccurate in my opinion. Instead, I would consider it a stable, completed project that no longer needs any active development. That said, there is a fork of the library created by the CoreOS team which can be found here: <https://github.com/coreos/bbolt>*

For now, don't worry about where you store the database that bolt connects to. At this stage I intend to just use whatever directory the `task` command was run from, so I will be using code roughly like this:

```go
db, err := bolt.Open("tasks.db", 0600, nil)
```

Later you can dig into how to install the application so that it can be run from anywhere and it will persist our tasks regardless of where we run the CLI.

### 3. Putting it all together

Finally, put the two pieces your wrote together so that when someone types `task add some task` it adds that task to the boltdb.

After that, explore how to setup and install the application so that it can be run from any directory in your terminal. This might require you to look into how to find a user's home directory on any OS (Windows, Mac OS, Linux, etc).

If you'd like, you can look into how to determine this on your own, but I recommend just grabbing this package: <https://github.com/mitchellh/go-homedir>. You can read over the code to see how it works - it is only 137 lines of code - but it should take care of all the oddities between different operating systems for us.

After that you will need to look into how to install a binary on your computer. The first place I suggest starting is the `go install` command. (*Hint: Try `go install --help` to see what this command does.*). This is likely to be the simplest route, but there are other options (like manually copying a binary to a directory in your `$PATH`).

*Note: I suspect many users will have issues around here that are OS specific. If you do, please first check the [Github issues](https://github.com/gophercises/task/issues?utf8=%E2%9C%93&q=is%3Aissue) to see if there are any open or closed issues that are similar to your problem. I'm hoping to use that as a nice Q&A section for this exercise.*

If all goes well you should have a complete CLI for managing your tasks installed once done with this section.


## Bonus

As a bonus exercise, I recommend working on the following two new commands:

```
$ task rm 1
You have deleted the "review talk proposal" task.

$ task completed
You have finished the following tasks today:
- wash the dishes
- clean the car
```

The `rm` command will delete a task instead of completing it.

The `completed` command will list out any tasks completed in the same day. You can define this however you want (last 12hrs, last 24hrs, or the same calendar date).

The first version of our CLI could get away with deleting tasks from the DB, but if you want these features to work you are likely going to need to tweak your DB design a bit. I'll leave that as an exercise for you to try out on your own, but if you need help feel free to get in touch - <jon@calhoun.io>

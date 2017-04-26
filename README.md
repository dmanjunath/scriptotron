# What is Script-O-tron?

It's a configurable script runner that can also notify you when you have errors or warnings. 

# Why?

Sometimes I need to execute long running scripts, but I dont want to be in front of my computer for hours at a time staring at the screen. At the same time, I don't want to have to litter my code with custom calls to raygun or rollbar or another monitoring service to be notified of crashes. I want a script runner which can do all of the above for me with no code changes whatsoever.

# How does it work?

Script-O-tron takes a configuration file config.toml which it parses for properties like command, logfile name and connection properties to gmail and twilio to send notifications. An example .toml file can be found at https://github.com/dmanjunath/script-O-tron/blob/master/example.toml

# How do I install and run it?

You can download the pre built packages from the releases section, or you can get the latest build from the bin/ folder. If you want to build the project yourself, just run 
```go
go build
```

To run, rename the example.toml to config.toml, and in the same folder level as config.toml run:

```shell
./bin/scriptotron
```

# What if Gmail isn't working?

In order to use gmail to send emails, you need to change the configuration to allow less secure apps here https://myaccount.google.com/lesssecureapps
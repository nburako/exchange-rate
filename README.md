# exchange-rate
This is a RESTful API that pulls & saves & converts the latest exchange rates from a third party API.



# DESIGN
I used MongoDB as NoSQL database and Gin Gonic as web framework. I also used Viper as a configuration solution.

There are 3 endpoints in the project - which are pulling data from third party API, saving all records into the database and converting rates according to the relevant parameters - with a Controller that handles them. There are Service and Repository layers used with the Dependency Injection pattern.

## Future Works

- [ ] Add SwaggerUI for documentation.
- [ ] Apply better panic/defer approaches.

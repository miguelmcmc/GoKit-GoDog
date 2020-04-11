Feature: create users
    In order to manage a company users
    I need to be able to create users

    Scenario: should crete users
        Given want to create this user
            | name   | email   |
            | <name> | <email> |
        When I executes create user function
        Then I list users
        Then the response should match <result>

        Examples:
            | name | email             | result                                          |
            | john | john.doe@mail.com | [{"name": "john","email": "john.doe@mail.com"}] |
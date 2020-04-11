Feature: create user
    In order to manage a company
    I need to be able to manage users

    Scenario: should list users
        Given there are users:
            | name | email             |
            | john | john.doe@mail.com |
            | jane | jane.doe@mail.com |
        When I list users
        Then the response should match:
            """
            [
                {
                    "name": "john",
                    "email": "john.doe@mail.com"
                },
                {
                    "name": "jane",
                    "email": "jane.doe@mail.com"
                }
            ]
            """
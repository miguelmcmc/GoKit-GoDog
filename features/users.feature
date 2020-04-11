Feature: users
    In order to manage a company
    I need to be able to manage users

    Scenario: should list no users when empty
        When I list users
        Then the response should match:
            """
            []
            """

    Scenario: should create user
        When I list users
        Then the response should match:
            """
            []
            """

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
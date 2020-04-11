Feature: list users
    In order to manage a company users
    I need to be able to list users

    Scenario: should list no users when empty
        When I list users
        Then the response should match:
            """
            []
            """

    Scenario: should list users
        Given there are users:
            | id | name | email             |
            | 0  | john | john.doe@mail.com |
            | 1  | jane | jane.doe@mail.com |
        When I list users
        Then the response should match:
            """
            [
                {
                    "id": "0",
                    "name": "john",
                    "email": "john.doe@mail.com"
                },
                {
                    "id": "1",
                    "name": "jane",
                    "email": "jane.doe@mail.com"
                }
            ]
            """
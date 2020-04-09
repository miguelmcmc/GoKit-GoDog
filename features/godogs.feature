Feature: eat godogs
    In order to be happy
    As a hungry gopher
    I need to be able to eat godogs

    Scenario: Eat 5 out of 12
        Given there are 12 godogs
        When I eat 5
        Then there should be 7 remaining

    Scenario: Eat 1 out of 999
        Given there are 999 godogs
        When I eat 1
        Then there should be 998 remaining

    Scenario: Eat 0 out of 999
        Given there are 999 godogs
        When I eat 0
        Then there should be 999 remaining




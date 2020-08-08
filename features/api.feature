Feature: API basics
  
  Scenario: Responds to healthchecks
    When I send "GET" request to "/health"
    Then the response code will be 200
    And the body will be "200 OK"

  
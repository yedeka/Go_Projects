This repository provides an implementation of the coding challenges procvided in the Coding Challenges blog [here](https://codingchallenges.fyi/challenges/intro/)

This is an attempt to solve the coding challenges for building simple tools as well as my personal goal of learning go while building these simple projects.

The repository contains free exercises provided on John Calhoun's site [here](https://courses.calhoun.io/courses/cor_gophercises)
Solutions for the problems attempted are as below 

1. [Quiz Game](https://courses.calhoun.io/lessons/les_goph_01) 
    - Builds a simple quiz game where users is asked for a CSV with simple questions and answers and then those questions are used to quiz the user.
    - **Concepts** : flags, GoRoutines, time and Channels   
2. [URL Shortner](https://courses.calhoun.io/lessons/les_goph_04) 
    - Builds a simple URL shortner to take in a shortened URL andget the complete URL for the suggested URL 
    - **Concepts** : YAML file parsing, introduction of MUX for routing
    - **Helper Packages** : net/http, gopkg.in/yaml.v3
3. [Choose Your Own Adventure(CYOA)-Web App](https://courses.calhoun.io/lessons/les_goph_06)
    - Build a minimalistic HTML template based we application that features a Simple HTML page with capabilitites of updating the page content based on configured JSON file. 
    - **Concepts** : GO HTML teamplate Rendering, JSON Parsing, Structs tagging for JSON fields, functional Options
    - **Helper Packages** : encoding/json, net/http
4. [Link Parser Utility](https://courses.calhoun.io/lessons/les_goph_16)
    - A simple utility that extracts all the links from a page given a HTML file
    - **Concepts** : HTML parsing
    - **Helper Packages**: golang.org/x/net/htm
5. [SiteMap builder](https://courses.calhoun.io/lessons/les_goph_24)
    - Utility built on top of the Link parser utility that takes in a seed Url and generates a site map of all the pages linked to the seed URL. Use BFS for getting links of all the child links obtianed from the seed link.
    - **Concepts** : Captrure resonse HTML of a GET request, defer, Copying data from reader to writer stream, BFS
    - **Helper Packages** : net/http  

This is the public source code repository of EFreedomWall - a finals project for the CC221 Web Development 2 course.

## Description
EFreedomWall is an application that allows visitors post their thoughts without creating an account. Anyone can like, update, and delete the post unless it is protected by a password. It is inspired from campus anonymous confession Facebook pages as well as other similar services such as CuriousCat.

## Architecture
EFreedomWall consists of three parts:

1. Web Service - the main component of the application. Where most of the operations and communication to the data storage happens. Uses WCF (particularly CoreWCF) as a requirement for this web service.
2. SOAP-to-REST API Service - another component that calls procedures on behalf of the client (aka the frontend) and converts the incoming SOAP response into JSON which the client will consume. It also acts as a frontend server in order to serve the frontend portion and the REST API service in one executable. Uses Go for this component.
3. Web App - A Single Page Application (SPA) that serves the user upon entering the website. Uses Vue and ViteJS as the frontend and build tool respectively.

## Notes
- While the second component can be merged with the web service, CoreWCF does not have this functionality compared to .NET WCF and recommends using ASP.NET Core directly which violates the requirement for this course.
- Used CoreWCF instead of WCF as I would need a Windows laptop to use Visual Studio (lol). Visual Studio for Mac has finite sets of features compared to the former.
  

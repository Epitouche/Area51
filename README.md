<div>
    <h1 align="center">SafeArea</h1>
</div>

<details>
    <summary>Table of Contents</summary>
    <ul>
        <li><a href="#about-the-project">About the Project</a></li>
        <li><a href="#main-features">Main Features</a></li>
        <li><a href="#project-architecture">Project Architecture</a></li>
        <li><a href="#getting-started">Getting Started</a>
            <ul>
                <li><a href="#prerequisites">Prerequisites</a></li>
                <li><a href="#installation">Installation</a></li>
                <li><a href="#usage">Usage</a></li>
            </ul>
        </li>
        <li><a href="#project-organization">Organization</a></li>
        <li><a href="#contributing">Contributing</a></li>
        <li><a href="#the-team">Team</a></li>
        <li><a href="#useful-links">Useful links</a></li>
    </ul>
</details>

## About the project
Welcome to the AREA Automation Suite, a powerful platform inspired by IFTTT and Zapier called the SafeAREA. This software allows users to automate workflows by connecting Actions and REActions through AREA configurations (Action-REAction-Automations). The suite includes the following components:

- Application Server: Manages core functionalities like user management, services, and AREA handling.
- Web Client: A browser-based interface to configure and monitor AREA workflows.
- Mobile Client: A mobile app for configuring and monitoring AREA workflows on the go.

## Main features
1. User Management:
    - Register, authenticate, and manage user accounts.
    - Confirm enrollment through email validation.
2. Services:
    - Subscribe to and configure services with Actions and REActions.
    - Examples: Webhooks, Email Notifications, Social Media Integrations.
3. AREA:
    - Create workflows by connecting an Action to a REAction.
    - Automatically trigger AREA workflows based on defined conditions.
4. Triggers:
    - Time-based and event-based triggers to automate workflows.

## Project Architecture
This project is divided into three main components:

1. Application Server:
    - Backend service handling all application logic.
    - Exposes a RESTful API for the clients.
    - Manages user authentication, AREA configurations, and triggers.
    - Built using Go, a simple but efficient language.
2. Web Client:
    - User-friendly browser interface for managing workflows.
    - Built using Nuxt.Js, a framework Vue.Js using Typescript and Tailwind CSS for easier styled components.
3. Mobile Client:
    - Mobile application for managing workflows on Android and iOS.
    - Provides push notifications and offline functionality.
    - Built in React Native for the simplicity of the mobile development.

## Getting Started
### Prerequisites
To set up the project, ensure you have the following tools installed:
- Vue.Js: VSCode extension or install it locally for web.
- Go: Go language for backend.
- Mobile Framework: React Native for mobile client.
- Dev Container: VSCode extension to easily launch backend/frontend.

### Installation
Clone the repository:
```bash
git clone git@github.com:Epitouche/Area51.git
```

### Usage

After cloning the repository go [here]() to follow the instruction to launch the backend or frontend in dev mode.

## Project Organization
We use the following tools for project management and documentation (Learn more by clicking on Github Projects or Google Drive):
- [GitHub Projects](./docs/Organization.md): For sprint planning, issue tracking, and task management. ([link](https://github.com/orgs/Epitouche/projects/2))
- [Google Drive](./docs/Organization.md): Central repository for sprint notes, design documents, and the proof of concept ([POC](./docs/POC.md)). ([link](https://drive.google.com/drive/folders/1Z0oZLYy2zBhhryj8Y1aOzdajEbtKuYpq))

## Contributing
We welcome contributions to improve this project !\
You can contribute by creating an issue on this project and the members of the team will take care of the problem or upgrade.

## The team
The team is composed of a group of 5 french students:
- [JsuisSayker](https://github.com/JsuisSayker)
- [OxiiLB](https://github.com/OxiiLB)
- [Dvaking](https://github.com/Dvaking)
- [Babouche](https://github.com/Babouuchee)
- [Karumapathetic](https://github.com/karumapathetic)

## Useful links
Here is the links to the markdown documentation for this project:
- [Organization](./docs/Organization.md)
- [Proof of Concepts (POC)](./docs/POC.md)
- [Services](./docs/Services.md)
- [Backend](./docs/Backend.md)
- [Frontend](./docs/Frontend.md)
- [Mobile](./docs/Mobile.md)

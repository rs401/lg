<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Go Report Card](https://goreportcard.com/badge/github.com/rs401/lg)](https://goreportcard.com/report/github.com/rs401/lg)
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">Go net/rpc Authentication system</h3>

  <p align="center">
    Go Microservices API with standard net/rpc, Docker containers and Kubernetes.
    <br />
    <a href="https://github.com/rs401/lg/docs"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/rs401/lg/issues">Report Bug</a>
    ·
    <a href="https://github.com/rs401/lg/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

I created this project as a way to learn more about Go's net/rpc, as well as learn more about Kubernetes and step up from docker-compose.

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Go](https://go.dev/)
* [Docker](https://www.docker.com/)
* [Kubernetes](https://kubernetes.io/)
* [mini-Kube](https://github.com/kubernetes/minikube)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

If you would like to test this project you can clone the repo and use the `make` to build and run. You will need Docker to build the docker images and minikube to run the kubernetes cluster.

### Prerequisites

You will need to edit the `env.sample` files and rename them to `.env`. You will also need a PostgreSQL database to connect to.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/rs401/lg.git
   ```
3. Build the auth service and client API
   ```sh
   make build-api
   make build-auth
   ```
4. To run locally you will need to:
   ```sh
   cd auth
   ./authsvc
   ```
   Open another terminal and:
   ```sh
   cd api
   ./authapi
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Once the services are running and connected to your db you can send a request to http://localhost:9000/api/signup/ with a json object with a name, email and password in the body.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/rs401/lg/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

I welcome any code reviews, tips, suggestions, etc.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

[![LinkedIn][linkedin-shield]][linkedin-url]

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-url]: https://github.com/rs401/lg/graphs/contributors
[forks-url]: https://github.com/rs401/lg/network/members
[stars-url]: https://github.com/rs401/lg/stargazers
[issues-url]: https://github.com/rs401/lg/issues
[license-url]: https://github.com/rs401/lg/blob/master/LICENSE.txt
[license-shield]: https://img.shields.io/github/license/rs4011/lg.svg?style=for-the-badge
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/richard-stadnick-3b4ab53b

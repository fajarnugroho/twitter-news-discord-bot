<!--
*** Thanks for checking out the twitter-news-discord-bot. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Twitter News Discord Bot</h3>

  <p align="center">
    A Discord Bot that follow tweets from twitter accounts and posted to a specific discord channel.
    <br />
    <a href="https://github.com/fajarnugroho/twitter-news-discord-bot"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/fajarnugroho/twitter-news-discord-bot/issues">Report Bug</a>
    ·
    <a href="https://github.com/fajarnugroho/twitter-news-discord-bot/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
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
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![Product Name Screen Shot][product-screenshot]

Twitter News Discord Bot is a personal project. This bot follow tweets from listed twitter accounts defined as `ACCOUNTS` in `.env` file. Tweets are obtained from Twitter Stream API and posted to Discord Channel which is defined as `DISCORD_CHANNEL_ID` in `.env` file.
Another feature will cooming soon including web dashboard, multi-channel post, tweets grouping, and many more.

### Built With

* [go-twitter](https://github.com/dghubble/go-twitter)
* [discordgo](https://github.com/bwmarrin/discordgo)


<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps.
### Prerequisites

* golang
* twitter API key
* discord bot token
* discord channel ID

### Installation

1. Get a twitter API Key at [https://developer.twitter.com](https://developer.twitter.com/en)
2. Create Discord Application and get bot token from [https://discord.com/developers/applications](https://discord.com/developers/applications)
3. Get Discord Channel ID from discord application
4. Clone the repo
   ```sh
   $ git clone https://github.com/fajarnugroho/twitter-news-discord-bot.git
   ```
5. Download go module
   ```sh
   $ go mod download
   ```
6. Build binary
   ```sh
   $ go build ./cmd/tndb -o bin/tndb
   ```
7. Copy `.env.example` to `.env` then set twitter api key, discord bot token, and discord channel id
8. Run bot
    ```sh
    $ ./tndb
    ```
<!-- USAGE EXAMPLES -->
## Usage

_For more examples, please refer to the [Documentation](docs/CONFIGURATION.md)_


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/fajarnugroho/twitter-news-discord-bot/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Your Name - [@nugrohof_](https://twitter.com/nugrohof_) - fajar@nugrohof.com

Project Link: [https://github.com/fajarnugroho/twitter-news-discord-bot](https://github.com/fajarnugroho/twitter-news-discord-bot)




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[issues-shield]: https://img.shields.io/github/issues/fajarnugroho/twitter-news-discord-bot?style=flat-square
[issues-url]: https://github.com/fajarnugroho/twitter-news-discord-bot/issues
[license-shield]: https://img.shields.io/github/license/fajarnugroho/twitter-news-discord-bot.svg?style=flat-square
[license-url]: https://github.com/fajarnugroho/twitter-news-discord-bot/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/fajar-nugroho
[product-screenshot]: images/screenshot.png

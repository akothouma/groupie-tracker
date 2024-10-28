- [GroupieTrackers](#groupietrackers)
- [Installation](#installation)
- [Usage](#usage)
- [Specificities](#specificities)
- [Dependencies](#dependencies)
- [Contributions](#contributions)
  - [License](#license)
  - [Authors](#authors)

# GroupieTrackers

Description
This project includes functionalities to fetch and display data from the GroupieTrackers API. It provides endpoints to retrieve information about artists and their locations. The project utilizes Go's net/http package for making HTTP requests, encoding/json for JSON encoding and decoding, and html/template for rendering HTML templates.

# Installation


1. **Install Go**: Ensure that you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

2. **Clone the repository**:
 ```bash
   git clone https://learn.zone01kisumu.ke/git/cliffootieno/ascii-art-reverse
   cd ascii-art-reverse
```Make sure to have an active internet connection to fetch data from the GroupieTrackers API.


3.**Run the program**: Use the following command to execute the program

```bash
    go run . 
```


# Usage

1. Start the server by running the program.

2. Open your web browser and navigate to http://localhost:8081 to view the list of artists.

# Specificities

Use the GetArtists endpoint to fetch and display information about artists.
Use the ArtistDetails endpoint to retrieve and display locations associated with a specific artist ID.
Endpoints
/artists: Fetches and displays information about artists.
/artistDetails?id={artist_id}: Retrieves and displays locations associated with a specific artist ID.

# Dependencies

net/http package for making HTTP requests
encoding/json package for JSON encoding and decoding
html/template package for rendering HTML templates

# Contributions

Contributions are welcome! Fork the repository, make your changes, and submit a pull request for review. Follow the project's coding standards and guidelines when contributing.

## License

This project is licensed under the **MIT** License. See the [LICENSE](LICENSE) file for more details.

## Authors

This program was built and maintained by

- [cliffootieno](https://learn.zone01kisumu.ke/git/cliffootieno)

  <img src="https://learn.zone01kisumu.ke/git/avatars/7c3793c3fac1a5908d1646d153555890?size=870" width="200">

* [wnjuguna](https://learn.zone01kisumu.ke/git/wnjuguna)

  <img src="https://learn.zone01kisumu.ke/git/avatars/c9b7b96426b4781d5a16fef462551fb5?size=870" width="200">

* [shfana](https://learn.zone01kisumu.ke/git/shfana)

  <img src="https://learn.zone01kisumu.ke/git/avatars/b82abc0b61d38ce3a680d3c04e2331c8?size=870" width="200">
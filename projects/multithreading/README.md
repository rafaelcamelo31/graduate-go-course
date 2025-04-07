# Multithreading

In this project, you will implement an API to get the first Response returned from two distinct endpoint.

Two requests are made simultaneously to following endpoints:

- `https://brasilapi.com.br/api/cep/v1/01153000 + cep`
- `http://viacep.com.br/ws/" + cep + "/json/`

# Requirements

- You must use the fastest response returned between the two and cancel slower one;
- The Response must be displayed in command line with datas of address and show which endpoint has sent;
- You must limit the Response time to 1 second, or timeout error must be shown.

# Testing

- Change directory to `graduate-go-course/projects/multithreading/cmd` then run `go run .`;
- Access your `localhost:8080`, then check your terminal for returned response.

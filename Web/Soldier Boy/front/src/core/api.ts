import axios from "axios";

const client = axios.create({
    baseURL: "https://hmlndr-api.onrender.com",
})

export default client;
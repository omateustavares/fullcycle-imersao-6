import axios from "axios";

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_NEST_HOST,
});

export default api;

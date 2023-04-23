import axios from "axios";

export async function useData() {
    const data = axios.get("/channel/get-data")
    console.log("[debug]data:", data)
}
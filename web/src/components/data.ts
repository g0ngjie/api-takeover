import axios from "axios";

export async function useData() {
    const { data } = await axios.get("/channel/get-data")
    console.log("[debug]data:", data)
    if (data.code === 100) return data.data || []
    return []
}
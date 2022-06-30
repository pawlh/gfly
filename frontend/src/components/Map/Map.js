import './Map.css'
import {CircleMarker, MapContainer, TileLayer} from 'react-leaflet'
import {useEffect, useState} from "react";

let fetched = false

export const Map = () => {

    const [markerData, setMarkerData] = useState([]);
    // const [fetched, setFetched] = useState(false)

    useEffect(() => {
        (async () => {
            if (!fetched) {
                fetched = true
                const response = await fetch('http://localhost:8080/api/points')
                setMarkerData(await response.json())
            }
        })()
    }, [])

    return (
        <div className="Map">
            <MapContainer id={'map'} center={[51.505, -0.09]} zoom={13} scrollWheelZoom={true} preferCanvas={true} infinite={false}>
                <TileLayer style={{filter: 'hue-rotate(45deg) !important'}}
                           attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                           url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                />

                {markerData?.map((point, index) => {
                    console.log(point)
                    return point.latitude && point.longitude ? <CircleMarker
                        key={index}
                        center={
                            [point.latitude, point.longitude]
                        }
                        radius={10}
                    /> : <></>
                })}
            </MapContainer>
        </div>
    )
}
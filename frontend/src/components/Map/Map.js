import './Map.css'
import {MapContainer, TileLayer} from 'react-leaflet'

export const Map = () => {
    return (
        <div className="Map">
            <MapContainer id={'map'} center={[51.505, -0.09]} zoom={13} scrollWheelZoom={true}>
                <TileLayer style={{filter: 'hue-rotate(45deg) !important'}}
                    attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                />
            </MapContainer>
        </div>
    )
}
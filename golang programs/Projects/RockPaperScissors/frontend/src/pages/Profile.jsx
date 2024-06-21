import React from 'react'
import "./global.css"

const Profile = () => {
  return (
    <div className='profile-page'>
        <div className='card'>
        <div>
            Profile Photo here
        </div>

        <div className='main'>
            <h1>Ketan Rathod</h1>
        </div>
        <button>Edit Profile</button>

        </div>

        <div className='games'>
            <h1>Previous Games Played</h1>
            <div>
                
            </div>
        </div>
    </div>
  )
}

export default Profile
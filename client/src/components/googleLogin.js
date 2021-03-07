import React from 'react';
import {GoogleLogin, GoogleLogout} from 'react-google-login';

const clientId = '757561505559-morsn6o57smovlvh0aohl64t7as1h4lp.apps.googleusercontent.com';

const onSignIn = (googleUser) => {
  console.log('onSignIn...')
  var profile = googleUser.getBasicProfile();
  console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
  console.log('Name: ' + profile.getName());
  console.log('Image URL: ' + profile.getImageUrl());
  console.log('Email: ' + profile.getEmail()); // This is null if the 'email' scope is not present.

  var id_token = googleUser.getAuthResponse().id_token;
  console.log('id_token', id_token)
}

const onLogout = (googleUser) => {
  console.log('onLogout...')
  console.log(googleUser)
}

export const GoogleLoginBtn = props => <GoogleLogin 
  clientId={clientId}
  onSuccess={onSignIn}
  onFailure={onSignIn}
  scope="profile email"
/>

export const GoogleLogoutBtn = props => <GoogleLogout 
  clientId={clientId}
  onLogoutSuccess={onLogout}
/>
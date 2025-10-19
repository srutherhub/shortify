export default function useIsLoggedIn() {
  const cookies = document.cookie.split("=");
  if (cookies[0] == "hanko") {
    return true;
  } else {
    return false;
  }
}

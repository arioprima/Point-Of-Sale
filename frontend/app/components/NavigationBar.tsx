import Link from "next/link"
import { FiSearch, FiShoppingCart, FiAlignJustify} from "react-icons/fi";

const NavigationBar = () => {
  const feather = require('feather-icons');
  return (
    <nav className="navbar">
      <Link href="/" className="navbar-logo">Anak<span>Sehat</span></Link>
      <div className="navbar-nav">
        <Link className="a" href="/home">Home</Link>
        <Link className="a" href="/about">About</Link>
        <Link className="a" href="/menu">Menu</Link>
        <Link className="a" href="/contact">Contact</Link>
      </div>   
      <div className="navbar-extra">
        <Link className="a" href="/login">Login</Link>
        <Link className="a" href="/register">Register</Link>
      </div>
    </nav>
  )
}

export default NavigationBar
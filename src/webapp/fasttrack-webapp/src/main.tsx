import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import {BrowserRouter, Route, Routes} from "react-router";

createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <BrowserRouter>
          <Routes>
              <Route path="/" element={<h1>Home Page</h1>}/>
              <Route path="/app" element={<App/>}/>
              <Route path="*" element={<h1>404 - Not Found</h1>} />
          </Routes>
      </BrowserRouter>
  </StrictMode>,
)

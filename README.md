# PTBE_ROKETIN

This documentation is made for testing API endpoints.

---

### ✅ Create Movie

**Method:** POST  
**URL:** `http://localhost:8080/movies`  
**Headers:** `Content-Type: application/json`

**Request Body (JSON):**
**Note : please try these examples one by one

Example 1
{
  "title": "The Great Adventure",
  "description": "A thrilling journey of heroes across unknown lands.",
  "duration": 120,
  "artists": ["John Doe", "Jane Smith"],
  "genres": ["Adventure", "Action"]
}

Example 2
{
  "title": "The Lost Journey",
  "description": "An epic adventure across forgotten lands.",
  "artists": ["Alice Smith", "John Doe"],
  "duration": 120,
  "genres": ["Adventure", "Fantasy"]
}

Example 3
{
  "title": "City Lights",
  "description": "A drama about the quiet lives of city dwellers.",
  "artists": ["John Doe", "Maria Garcia"],
  "duration": 95,
  "genres": ["Drama"]
}

Example 4
{
  "title": "Space Runners",
  "description": "A sci-fi tale of smugglers racing through galaxies.",
  "artists": ["Liam Wong", "Nina Patel"],
  "duration": 110,
  "genres": ["Sci-Fi", "Action"]
}

**Response Example:**
{"id":1,"title":"The Great Adventure","description":"A thrilling journey of heroes across unknown lands.","duration":120,"artists":["John Doe","Jane Smith"],"genres":["Adventure","Action"]}

---

### 🔄 Update Movie

**Method:** PUT  
**URL:** `http://localhost:8080/movies?ID=2`  
**Headers:** `Content-Type: application/json`

**Request Body (JSON):**
{  
 "title": "La La Land",
  "description": "This is La La Land Movies",
  "duration": 150,
  "artists": ["Ryan “Gosling", "Emma Watson"],
  "genres": ["Romance”, “Musical"]
}

**Response Example:**
{
    "id": 2,
    "title": "La La Land",
    "description": "This is La La Land Movies",
    "duration": 150,
    "artists": [
        "Ryan “Gosling",
        "Emma Watson"
    ],
    "genres": [
        "Romance”, “Musical"
    ]
}

---

### 🔍 Search Movies

**Method:** GET  
**URL:** `http://localhost:8080/movies/search?q=John`  
**Headers:** `Content-Type: application/json`

**Response Example:**
[
    {
        "id": 1,
        "title": "The Great Adventure",
        "description": "A thrilling journey of heroes across unknown lands.",
        "duration": 120,
        "artists": [
            "John Doe",
            "Jane Smith"
        ],
        "genres": [
            "Adventure",
            "Action"
        ]
    },
    {
        "id": 3,
        "title": "City Lights",
        "description": "A drama about the quiet lives of city dwellers.",
        "duration": 95,
        "artists": [
            "John Doe",
            "Maria Garcia"
        ],
        "genres": [
            "Drama"
        ]
    }
]

---

### 📄 List Movies with Pagination

**Method:** GET  
**URL:** `http://localhost:8080/movies?page=1&limit=2`  
**Headers:** `Content-Type: application/json`

**Response Example:**
[
    {
        "id": 1,
        "title": "The Great Adventure",
        "description": "A thrilling journey of heroes across unknown lands.",
        "duration": 120,
        "artists": [
            "John Doe",
            "Jane Smith"
        ],
        "genres": [
            "Adventure",
            "Action"
        ]
    },
    {
        "id": 2,
        "title": "La La Land",
        "description": "This is La La Land Movies",
        "duration": 150,
        "artists": [
            "Ryan “Gosling",
            "Emma Watson"
        ],
        "genres": [
            "Romance”, “Musical"
        ]
    }
]






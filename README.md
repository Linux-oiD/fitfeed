# 🏋️ FitFeed: Privacy-First Fitness Social Platform

## 🌟 About FitFeed

FitFeed is a **self-hosted, privacy-first** social platform designed for fitness enthusiasts.

Think of it as your personal, open-source analogue to platforms like Strava. Users can share a feed of their fitness activities, view updates from friends, and interact with posts through likes and comments, all within a secure and controlled environment.

---

## ✨ Features

* **Activity Feed:** Share and view fitness activities with friends.
* **Self-Hosted:** Full control over your data and infrastructure.
* **Privacy-First:** Designed with user privacy as the top priority.

### 🚀 Future Plans (Roadmap)

I'm actively planning to expand FitFeed's capabilities to include:

* Social Interaction: likes and comments.
* Synchronization services for popular fitness devices (Garmin, Polar, Suunto).
* Integration with mobile health platforms (Google Fit, Apple Health).
* Storage and tracking of basic health measurements (weight, sleep, calories burned, blood pressure).

---

## 🛠️ Technology Stack

FitFeed is built as a set of microservices utilizing modern and efficient technologies:

| Component | Technology | Description |
| :--- | :--- | :--- |
| **Backend** | Go (Golang) | High-performance, compiled backend services. |
| **Frontend** | TypeScript, React, Ant Design | Robust, type-safe user interface built with the Ant Design component library. |
| **Build/Bundler** | Vite, Bun | Fast development server and high-performance package management/runtime. |
| **Database ORM/Migration** | Gorm, Goose | Used for database interaction and managing schema evolution. |

### 📁 Project Structure (Services)

The core functionality is split across several services located in the `services/` directory:

* `auth`: Handles user registration and authentication logic.
* `api`: The main application programming interface service.
* `dbm`: A custom database migration tool leveraging Goose and Gorm to manage database schema updates.
* `web`: The frontend client application (React + Ant Design).

---

## ⚙️ Installation and Setup

*(Note: Please fill in the detailed steps here once the setup is finalized. For now, this serves as a placeholder.)*

1.  **Prerequisites:** Ensure you have [Go/Bun/Docker/etc.] installed.
*TBD*
2.  **Clone the Repository:**
    ```bash
    git clone [https://github.com/Linux-oiD/fitfeed.git](https://github.com/Linux-oiD/fitfeed.git)
    cd fitfeed
    ```
3.  **Database Setup:** [Instructions to set up the database, e.g., PostgreSQL].
*TBD*
4.  **Run Migrations:**
    ```bash
    ./services/dbm migrate up 
    ```
5.  **Build and Run Services:**
    ```bash
    # Example commands to start backend and frontend
    # ...
    ```

---

## 📜 License

This project is licensed under the **AGPL (GNU Affero General Public License)**. See the `LICENSE` file for more details.

---

## 🔗 Live Demo / Documentation

* **Website:** [fitfeed.org](http://fitfeed.org/)

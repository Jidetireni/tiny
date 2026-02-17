# System Design: URL Shortener Service

## 🛠️ Non-Functional Requirements

### 1. Low Latency (The "Need for Speed")

Redirection must be perceived as "real-time." While a 200ms goal is standard, top-tier systems aim for **<50ms**.

* **Redirect Strategy:** Use **HTTP 302 (Found)** instead of 301. While 301 reduces server load via browser caching, 302 ensures every click hits our system, allowing for accurate **analytics and tracking**.
* **Caching Layer:** Implement a **Distributed Cache (Redis)** using an **LRU (Least Recently Used)** eviction policy to store the most frequently accessed mappings.
* **Read/Write Splitting:** Given the **100:1 read-to-write ratio**, the read path is isolated and hyper-optimized.

### 2. High Availability & Fault Tolerance

The system must remain operational even during regional outages or malicious attacks.

* **Database Replication:** Use **Cassandra (NoSQL)** for its masterless architecture, ensuring no single point of failure. Deploy in **Multi-Region** setups (e.g., US-East and EU-West) for geographic redundancy.
* **Rate Limiting:** Implement strict rate limiting at the **API Gateway** level to prevent ID exhaustion and DDoS attacks.
* **CAP Trade-off:** In the event of a network partition, the system prioritizes **Availability over Consistency** (AP), ensuring users can still be redirected even if the absolute latest URL update hasn't propagated to every node yet.

### 3. Scalability & Storage

The system must handle massive growth in both request volume and data footprint.

* **Horizontal Scaling:** Use **Docker/Kubernetes** to scale stateless application instances based on CPU/Memory load.
* **Data Retention (TTL):** Implement an **Expiration Policy**. To prevent "zombie" data from bloating the DB, links should expire (e.g., after 5 years) unless renewed.
* **Capacity Estimation:** At 100M links/month, the system requires ~6TB of storage over 5 years. Cassandra's column-oriented nature scales efficiently for this volume.

### 4. Uniqueness & Security

* **ID Generation:** Use **Zookeeper** to manage distributed ID ranges. This prevents collisions without requiring a centralized "auto-increment" lock on the database.
* **Base62 Encoding:** IDs are encoded using . A 7-character string provides  (~3.5 trillion) unique combinations.
* **Anti-Scraping:** To prevent users from guessing sequential URLs, avoid simple increments. Use a **shuffling algorithm** or a random component in the ID generation process to ensure non-predictability.

---

## 📊 Summary of Architectural Choices

| Component | Choice | Reason |
| --- | --- | --- |
| **Database** | Cassandra | High write throughput and seamless horizontal scaling. |
| **Caching** | Redis (LRU) | Sub-millisecond lookups for "hot" links. |
| **ID Generation** | Zookeeper + Ranges | High-performance uniqueness in a distributed environment. |
| **Protocol** | 302 Redirect | Required for real-time telemetry and analytics. |
| **Load Balancing** | API Gateway | Distributes traffic and handles global rate limiting. |

---

## 🔍 Future Considerations

* **Analytics Pipeline:** Integration with Kafka and Spark to track user demographics and click patterns.
* **Safe Browsing:** Integration with Google Safe Browsing API to black-list malicious destination URLs.
* **Custom Aliases:** Support for user-defined slugs (e.g., `short.com/my-link`).

---

**Would you like me to calculate the specific "back-of-the-envelope" math for your daily bandwidth and storage requirements?**

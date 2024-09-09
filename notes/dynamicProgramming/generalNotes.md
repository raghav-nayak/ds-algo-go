Dynamic programming
- memoization
- tabulation

Height of the tree -> distance between the root node and the farther descendent


memoization recipe
1. make it work -> brute force, without memoization
	- visualize the problem as a tree
	- implement tree using recursion
	- test it

2. make it efficient -> add memoization
	- add a memo object
	- add a base case to return memo values
	- store return values into the memo


![[Pasted image 20240903222518.png]]


![[Pasted image 20240903222534.png]]


![[Pasted image 20240903222600.png]]


Dynamic programming (DP) is a powerful technique used to solve complex problems by breaking them down into simpler subproblems and solving each subproblem just once, storing the results for future reference. This approach is widely applicable in various real-life scenarios, including:

### 1. **Route Optimization (Traveling Salesman Problem)**
   - **Scenario**: Delivery services like FedEx or UPS need to determine the most efficient route for delivering packages to multiple locations.
   - **Usage**: Dynamic programming helps optimize the route to minimize travel time or distance by storing and reusing results of previously computed routes, reducing redundant calculations.

### 2. **Resource Allocation**
   - **Scenario**: Companies need to allocate limited resources (like budget, time, or workforce) across different projects to maximize profit or efficiency.
   - **Usage**: DP can be used to solve resource allocation problems by breaking down the problem into stages, each representing a decision on resource distribution, and optimizing for the best overall outcome.

### 3. **Financial Modeling (Knapsack Problem)**
   - **Scenario**: Investors want to maximize returns while adhering to a budget constraint.
   - **Usage**: DP is used in portfolio optimization to select a combination of investments that maximizes returns while staying within budget, similar to the knapsack problem where items with different values and weights must be packed into a limited-capacity knapsack.

### 4. **Predictive Text and Autocorrect**
   - **Scenario**: Smartphones and text editors suggest words or correct spelling mistakes as users type.
   - **Usage**: DP is used in algorithms like the Levenshtein distance to calculate the minimum number of edits (insertions, deletions, or substitutions) required to transform one word into another, aiding in spell-checking and predictive text.

### 5. **Bioinformatics**
   - **Scenario**: Scientists compare DNA, RNA, or protein sequences to understand evolutionary relationships or identify genes.
   - **Usage**: DP algorithms like the Needleman-Wunsch or Smith-Waterman algorithms are used to align sequences optimally by storing and reusing solutions to subproblems, which helps in identifying similarities and differences between biological sequences.

### 6. **Inventory Management**
   - **Scenario**: Retailers need to decide how much stock to order to meet demand without overstocking.
   - **Usage**: DP helps in solving the inventory management problem by determining the optimal order quantity and timing to minimize costs while meeting demand, considering factors like holding costs, ordering costs, and shortage costs.

### 7. **Speech Recognition**
   - **Scenario**: Voice-activated systems, like virtual assistants (e.g., Siri, Google Assistant), convert spoken language into text.
   - **Usage**: DP algorithms like the Viterbi algorithm are used in Hidden Markov Models (HMMs) to determine the most likely sequence of words or phonemes given the input speech, enhancing accuracy in speech recognition.

### 8. **Software Development (Caching and Memoization)**
   - **Scenario**: Developers need to optimize applications to run efficiently, especially when dealing with repetitive tasks.
   - **Usage**: DP is applied in software development through caching (storing results of expensive function calls) and memoization (storing results of recursive function calls) to avoid redundant calculations and improve performance.

### 9. **Supply Chain Optimization**
   - **Scenario**: Businesses need to manage the flow of goods from suppliers to customers efficiently.
   - **Usage**: DP helps in solving supply chain optimization problems, such as determining the best way to transport goods across different routes and warehouses to minimize costs and meet delivery deadlines.

### 10. **Machine Learning (Reinforcement Learning)**
   - **Scenario**: AI systems learn to make decisions through trial and error to maximize rewards, such as in game playing or autonomous driving.
   - **Usage**: DP techniques like value iteration and policy iteration are used in reinforcement learning to compute the optimal policy by evaluating and improving the decisions made by the AI system at each step.

Dynamic programming’s ability to efficiently solve problems involving overlapping subproblems and optimal substructure makes it a valuable tool in these and many other real-world applications.
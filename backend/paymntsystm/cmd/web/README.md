**cmd/web/main.go What This Code Does**

This Go code creates a web application.  Imagine a website like an online store or a simple blog. This code sets up the foundation for how that website runs.

**The Thought Process: How to Write It**

1. **The Building Blocks (Structs)**
   * `config`: This holds the settings for your web app. It stores things like:
     * Which TCP/IP port to run on (like channel 4000 on a TV)
     * Whether you're in "development mode" (for testing) or "production mode" (for the live website)
     * Where to find another important part of your system, the API (more on that later)
     * Secrets for interacting with a payment provider (like Stripe)

   * `application`: This represents your whole running web app. It has:
     * The `config` we just talked about
     * Loggers to keep track of what's happening ("info" and "error" logs)
     * A place to store pre-made templates (like cookie-cutter designs for web pages)
     * The version number of the app

2. **Starting the Server (The `serve` function)**
   * This is like opening the doors to your web store. It tells your computer:
     * "Listen on this port (specified in the config)."
     * "Use these special routes (we'll see those later) to handle requests."
     * "Set some timeouts so things don't hang forever."

3. **The Main Event (`main` function)**
   * This is where the action begins:
      * It reads the command-line flags (e.g., `-port 8080` to change the port).
      * It pulls secrets (like your Stripe key) from environment variables (special places on your computer).
      * It creates loggers.
      * It sets up the `application` struct with all these pieces.
      * It calls the `serve` function to start the web server.

**How It All Works Together**

1. You run the code.
2. It reads settings and gets ready.
3. It starts listening for requests (people visiting your website).
4. When someone visits a page:
   * It finds the right route (like `/products` or `/blog`).
   * It might fetch data from the API.
   * It fills in a template to create the HTML for the page.
   * It sends that HTML back to the visitor's browser.

**Understanding APIs**
An API (Application Programming Interface) is like a waiter in a restaurant. Your website (the customer) asks for data, and the API (the waiter) goes to the kitchen (the database or other storage), gets the data, and brings it back to your website. This code assumes you have a separate API running somewhere else.

**Example: A Blog Post**

Let's say someone visits `/blog/my-first-post`. Here's a simplified version of what might happen:

1. The web server sees the request.
2. It checks the routes and finds a match.
3. It calls the API and asks for the blog post data.
4. The API gets the post from the database.
5. The web server fills in a template (like "blog post layout") with the post's title, content, etc.
6. It sends the final HTML to the browser, and you see the blog post!

****

**cmb/web/router.go What This Code Does**

This code handles the way your web application creates and displays web pages to users. Imagine you have a template for a product page, a blog post, or a checkout form. This code figures out how to take those templates and fill them with the right data to show the user.

**Thought Process: How to Write It**

1. **Data Structure (`templateData`)**
   * Think of `templateData` as a container to hold all the information you want to display on a web page. It can store:
     * Strings (text like product names, titles)
     * Integers (numbers like prices, quantities)
     * Floating-point numbers (numbers with decimals, maybe for discounts)
     * General data (`interface{}` can hold different types)
     * Security token (CSRFToken, helps protect against attacks)
     * Messages for the user (flash, warning, error)
     * Whether the user is logged in
     * API URL (so templates can know where to fetch more data)
     * Version of the stylesheet (CSS)

2. **Functions and Embedded Templates (`functions`, `templateFS`)**
   * `functions` is a way to add custom functions that you can use directly in your templates (for example, to format dates or prices nicely).
   * `templateFS` (using `//go:embed`) embeds the actual template files (like HTML files) into your Go code. This makes it easier to package and distribute your application.

3. **Adding Default Data (`addDefaultData`)**
   * This function adds standard information to the `templateData` before a page is rendered. You might add things like the current year, the website name, or links to common pages.

4. **Rendering Templates (`renderTemplate`)**
   * This is the core function:
     * It figures out which template file to use based on the `page` name.
     * If you're in production mode and the template has already been parsed, it reuses it from a cache (`templateCache`) to save time.
     * Otherwise, it parses the template and its partials (smaller pieces of templates) from the embedded file system.
     * It adds any default data.
     * Finally, it "executes" the template, filling it with the data and sending the resulting HTML to the user's web browser.

5. **Parsing Templates (`parseTemplate`)**
   * This helper function takes care of reading the template files, combining the main template with its partials, and making it ready for rendering. It then stores the parsed template in the cache for later use.

**How It All Works Together**

1. When the web server in the previous code receives a request (like for `/products/123`), it figures out which page and route to use.
2. It calls this `renderTemplate` function, providing the page name and any data needed for that page.
3. This code finds the right template, prepares it, fills it with data, and sends the completed web page back to the user.

**Example: A Product Page**

Imagine a request to `/products/123`. Here's a simplified version of what might happen:

1. The web server calls `renderTemplate("product", td)`, where `td` contains the product details.
2. The code finds the `templates/product.page.tmpl` file.
3. It might include partials like `templates/header.partial.tmpl` and `templates/footer.partial.tmpl`.
4. It fills in the template with data from `td`: product name, price, description, etc.
5. It sends the final HTML to the browser, and you see the product page!

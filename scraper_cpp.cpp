// #include <iostream>
// #include <curl/curl.h>
// #include <regex>

// int main() {
//     CURL * curl = curl_easy_init();
//     std::string response;
//     curl_easy_setopt(curl, CURLOPT_URL, "https://books.toscrape.com/");
//     curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, [](void* data, size_t n, std::string* out)
//     { out->append((char*)data, s*n); 
//         return s*n;});
//     curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);
//     std::regex r("product_pod");
//     int count = std::distance(std::sregex_iterator(response.begin(), response.end(), r), std::sregex_iterator());
//     std::cout << count;
//     return 0;
// }
// copier oller de l'ia
#include <iostream>
#include <curl/curl.h>
#include <regex>

size_t write_callback(void* data, size_t size, size_t nmemb, std::string* out) {
    out->append((char*)data, size * nmemb);
    return size * nmemb;
}

int main() {
    CURL* curl = curl_easy_init();
    if (!curl) return 1;

    std::string response;
    curl_easy_setopt(curl, CURLOPT_URL, "https://books.toscrape.com/");
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);
    curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1L);
    
    curl_easy_perform(curl);
    curl_easy_cleanup(curl);

    std::regex r("product_pod");
    int count = std::distance(
        std::sregex_iterator(response.begin(), response.end(), r),
        std::sregex_iterator()
    );
    
    std::cout << count;
    return 0;
}


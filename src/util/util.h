#pragma once

#include <concepts>
#include <span>

template <typename T>
requires std::equality_comparable<T>
bool contains(std::span<const T> arr, const T& el) {
    for(const auto& e: arr) {
        if (e == el) return true;
    }
    return false;
}
